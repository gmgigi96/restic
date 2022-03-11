//go:build darwin || freebsd || linux
// +build darwin freebsd linux

package fuse

import (
	"context"

	"bazil.org/fuse"
	"github.com/restic/restic/internal/restic"
)

type other struct {
	root *Root
	node *restic.Node
}

func newOther(ctx context.Context, root *Root, node *restic.Node) (*other, error) {
	return &other{root: root, node: node}, nil
}

func (l *other) Readlink(ctx context.Context, req *fuse.ReadlinkRequest) (string, error) {
	return l.node.LinkTarget, nil
}

func (l *other) Attr(ctx context.Context, a *fuse.Attr) error {
	a.Inode = l.node.Inode
	a.Mode = l.node.Mode

	if !l.root.cfg.OwnerIsRoot {
		a.Uid = l.node.UID
		a.Gid = l.node.GID
	}
	a.Atime = l.node.AccessTime
	a.Ctime = l.node.ChangeTime
	a.Mtime = l.node.ModTime

	a.Nlink = uint32(l.node.Links)

	return nil
}
