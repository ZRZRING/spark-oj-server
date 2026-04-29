package sshtunnel

import (
	"fmt"
	"io"
	"net"
	"os"
	"path/filepath"
	"sync"

	"github.com/gogf/gf/v2/frame/g"
	"github.com/gogf/gf/v2/os/gctx"
	"golang.org/x/crypto/ssh"
)

type Forward struct {
	Local  string `json:"local"`
	Remote string `json:"remote"`
}

type Config struct {
	Host       string    `json:"host"`
	User       string    `json:"user"`
	PrivateKey string    `json:"privateKey"`
	Forwards   []Forward `json:"forwards"`
}

func loadConfig() (*Config, error) {
	ctx := gctx.GetInitCtx()
	cfg := g.Cfg()

	host := cfg.MustGet(ctx, "sshTunnel.host").String()
	if host == "" {
		return nil, nil
	}

	user := cfg.MustGet(ctx, "sshTunnel.user").String()
	keyPath := cfg.MustGet(ctx, "sshTunnel.privateKey").String()

	var forwards []Forward
	err := cfg.MustGet(ctx, "sshTunnel.forwards").Scan(&forwards)
	if err != nil {
		return nil, fmt.Errorf("parse sshTunnel.forwards: %w", err)
	}

	return &Config{
		Host:       host,
		User:       user,
		PrivateKey: keyPath,
		Forwards:   forwards,
	}, nil
}

func expandPath(path string) string {
	if len(path) > 0 && path[0] == '~' {
		home, _ := os.UserHomeDir()
		return filepath.Join(home, path[1:])
	}
	return path
}

func connectSSH(cfg *Config) (*ssh.Client, error) {
	keyPath := expandPath(cfg.PrivateKey)
	keyData, err := os.ReadFile(keyPath)
	if err != nil {
		return nil, fmt.Errorf("read private key %s: %w", keyPath, err)
	}

	signer, err := ssh.ParsePrivateKey(keyData)
	if err != nil {
		return nil, fmt.Errorf("parse private key: %w", err)
	}

	client, err := ssh.Dial("tcp", cfg.Host, &ssh.ClientConfig{
		User: cfg.User,
		Auth: []ssh.AuthMethod{ssh.PublicKeys(signer)},
		// #nosec G106
		HostKeyCallback: ssh.InsecureIgnoreHostKey(),
	})
	if err != nil {
		return nil, fmt.Errorf("ssh dial %s: %w", cfg.Host, err)
	}

	return client, nil
}

func forwardOnce(client *ssh.Client, fwd Forward) error {
	localAddr := fwd.Local
	remoteAddr := fwd.Remote

	listener, err := net.Listen("tcp", localAddr)
	if err != nil {
		return fmt.Errorf("listen %s: %w", localAddr, err)
	}

	go func() {
		for {
			local, err := listener.Accept()
			if err != nil {
				return
			}

			go func() {
				remote, err := client.Dial("tcp", remoteAddr)
				if err != nil {
					local.Close()
					return
				}

				var wg sync.WaitGroup
				wg.Add(2)
				go func() {
					defer wg.Done()
					io.Copy(remote, local)
					remote.Close()
				}()
				go func() {
					defer wg.Done()
					io.Copy(local, remote)
					local.Close()
				}()
				wg.Wait()
			}()
		}
	}()

	return nil
}

func Setup() error {
	cfg, err := loadConfig()
	if err != nil {
		return fmt.Errorf("load sshTunnel config: %w", err)
	}
	if cfg == nil {
		g.Log().Info(gctx.GetInitCtx(), "sshTunnel not configured, skipping")
		return nil
	}

	client, err := connectSSH(cfg)
	if err != nil {
		return err
	}

	for _, fwd := range cfg.Forwards {
		if err := forwardOnce(client, fwd); err != nil {
			client.Close()
			return err
		}
		g.Log().Infof(gctx.GetInitCtx(), "ssh tunnel: %s -> %s", fwd.Local, fwd.Remote)
	}

	g.Log().Infof(gctx.GetInitCtx(), "ssh tunnel established via %s", cfg.Host)
	return nil
}
