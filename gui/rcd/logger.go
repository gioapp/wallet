package rcd

var (
	// MaxLogLength is a var so it can be changed dynamically
	MaxLogLength = 16384
)

func (r *RcVar) DuoUIloggerController() {
	//logChan := logi.L.AddLogChan()
	//logi.L.SetLevel(*r.cx.Config.LogLevel, true, "pod")
	//go func(gtx layout.Context)layout.Dimensions{
	//out:
	//	for {
	//		select {
	//		case n := <-logChan:
	//			le, ok := r.Log.LogMessages.Load().([]logi.Entry)
	//			if ok {
	//				le = append(le, n)
	//				// Once length exceeds MaxLogLength we trim off the start to keep it the same size
	//				ll := len(le)
	//				if ll > MaxLogLength {
	//					le = le[ll-MaxLogLength:]
	//				}
	//				r.Log.LogMessages.Store(le)
	//			} else {
	//				r.Log.LogMessages.Store([]logi.Entry{n})
	//			}
	//		case <-r.Log.StopLogger:
	//			defer func(gtx layout.Context)layout.Dimensions{
	//				r.Log.StopLogger = make(chan struct{})
	//			}()
	//			r.Log.LogMessages.Store([]logi.Entry{})
	//			logi.L.LogChan = nil
	//			break out
	//		}
	//	}
	//	close(logChan)
	//}()
}
