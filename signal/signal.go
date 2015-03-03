// Copyright (c) 2014 Xiaomi.com, Inc. All Rights Reserved
// @file    signal.go
// @author  王靖 (wangjing1@xiaomi.com)
// @date    14-11-26 15:36:24
// @version $Revision: 1.0 $
// @brief

package signal2

import (
	. "github.com/wangjild/go-mysql-proxy/log"
	"os"
)

type SignalHandler func(s os.Signal, arg interface{}) error

type SignalSet struct {
	M map[os.Signal]SignalHandler
}

func NewSignalSet() *SignalSet {
	s := new(SignalSet)
	s.M = make(map[os.Signal]SignalHandler)
	return s
}

func (s *SignalSet) Register(sig os.Signal, handler SignalHandler) {
	if _, exist := s.M[sig]; !exist {
		s.M[sig] = handler
	}
}

func (s *SignalSet) Handle(sig os.Signal, arg interface{}) error {
	if handler, exist := s.M[sig]; exist {
		return handler(sig, arg)
	} else {
		SysLog.Warn("no available handler for signal %v, ignore!", sig)
		return nil
	}
}

func init() {

}

/* vim: set expandtab ts=4 sw=4 */
