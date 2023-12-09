package events

import "cabbage-server/boot"

func init() {
	// the comment create event
	boot.Emit.On("comment:create", func(args ...any) {
		boot.Logger.Info("the comment create")
	})
	// the comment delete event
	boot.Emit.On("comment:delete", func(args ...any) {
		boot.Logger.Info("the comment delete")
	})
	// the comment update event
	boot.Emit.On("comment:update", func(args ...any) {
		boot.Logger.Info("the comment update")
	})
	// the comment operator event
	boot.Emit.On("comment:operator", func(args ...any) {
		boot.Logger.Info("the comment operater")
	})
	// the post create event
	boot.Emit.On("post:create", func(args ...any) {
		boot.Logger.Info("the post create")
	})
	// the post delete event
	boot.Emit.On("post:delete", func(args ...any) {

		boot.Logger.Info("the post delete")
	})
	// the post update event
	boot.Emit.On("post:update", func(args ...any) {

		boot.Logger.Info("the post update")
	})
	// the post search event
	boot.Emit.On("post:search", func(args ...any) {

		boot.Logger.Info("the post search")
	})
	// the post operator event
	boot.Emit.On("post:operator", func(args ...any) {
		boot.Logger.Info("the post operator")

	})

	// the post operator event
	boot.Emit.On("tag:create", func(args ...any) {

		boot.Logger.Info("the tag create")
	})

	boot.Emit.On("tag:delete", func(args ...any) {
		boot.Logger.Info("the tag delete")

	})

	// the post operator event
	boot.Emit.On("user:login", func(args ...any) {
		boot.Logger.Info("the post create")
	})

	boot.Emit.On("user:register", func(args ...any) {

		boot.Logger.Info("the user register")
	})

	boot.Emit.On("user:ban", func(args ...any) {

		boot.Logger.Info("the user is ban")
	})

	boot.Emit.On("user:follow", func(args ...any) {

		boot.Logger.Info("the user follow")
	})

	boot.Emit.On("user:unfollow", func(args ...any) {
		boot.Logger.Info("the user unfollow")

	})
}
