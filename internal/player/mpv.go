package player

import (
	"os/exec"
)

type MPVPlayer struct {
	cmd   *exec.Cmd
	state State
}

func NewMPVPlayer() *MPVPlayer {
	return &MPVPlayer{
		state: StateStopped,
	}
}

func (p *MPVPlayer) Play(url string) error {
	_ = p.Stop() // Ensure any existing instance is stopped

	p.cmd = exec.Command("mpv", "--no-video", "--really-quiet", url)
	err := p.cmd.Start()
	if err != nil {
		p.state = StateError
		return err
	}

	p.state = StatePlaying

	// Capture command locally to avoid race condition with p.Stop() setting p.cmd to nil
	cmd := p.cmd
	// Wait in background to reset state when process ends
	go func() {
		if cmd != nil {
			_ = cmd.Wait()
		}
		p.state = StateStopped
	}()

	return nil
}

func (p *MPVPlayer) Stop() error {
	if p.cmd != nil && p.cmd.Process != nil {
		err := p.cmd.Process.Kill()
		if err != nil {
			return err
		}
		p.cmd = nil
	}
	p.state = StateStopped
	return nil
}

func (p *MPVPlayer) State() State {
	return p.state
}
