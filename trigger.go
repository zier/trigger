/*
* @Author: Xier
* @Date:   2015-02-04 13:00:49
* @Last Modified by:   Xier
* @Last Modified time: 2015-02-04 14:35:40
 */

package trigger

import (
	"time"
)

type Trigger struct {
	Timing time.Time
}

func Start() *Trigger {
	return &Trigger{Timing: time.Now()}
}
func (tr *Trigger) Stop() time.Duration {
	return time.Now().Sub(tr.Timing)
}
func (tr *Trigger) Reset() {
	tr.Timing = time.Now()
}

// tr := trigger.Start()
//   trtime := tr.Stop()
//   if trtime > 1000 {
//     fmt.Println(trtime.Nanoseconds())
//   }
