package main

import (
	bst "./BinarySearchTree"

	"go.uber.org/zap"
)

func initlogger() *zap.Logger {
	config := zap.NewDevelopmentConfig()
	logger, _ := config.Build()
	return logger
}

func main() {
	loggerMgr := initlogger()
	zap.ReplaceGlobals(loggerMgr)
	defer loggerMgr.Sync() // flushes buffer, if any
	logger := loggerMgr
	logger.Debug("START!!!!")
	logger.Info("Binary Search Tree")
	tree := &bst.BST{}
	//calling inorder on empty tree
	tree.Inorder()
	//adding nodes into the bst
	tree.Insert(20)
	tree.Insert(15)
	tree.Insert(10)
	tree.Insert(28)
	tree.Insert(27)
	//inorder traversal of the bst
	tree.Inorder()
}
