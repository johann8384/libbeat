package stdout

import (
  "time"

  "github.com/johann8384/libbeat/common"
  "github.com/johann8384/libbeat/logp"
  "github.com/johann8384/libbeat/outputs"
)

type StdOutput struct {
  enabled string
}


func (out *StdOutput) Init(config outputs.MothershipConfig, topology_expire int) error {
  // not supported by this output type
  return nil
}

func (out *StdOutput) PublishIPs(name string, localAddrs []string) error {
  // not supported by this output type
  return nil
}

func (out *StdOutput) GetNameByIP(ip string) string {
  // not supported by this output type
  return ""
}

func (out *StdOutput) PublishEvent(ts time.Time, event common.MapStr) error {
  //json_event, err := json.Marshal(event)
  //if err != nil {
  //  logp.Err("Fail to convert the event to JSON: %s", err)
  //  return err
  //}

  out.Print(event)
  return nil
}

func (out *StdOutput) Print(event common.MapStr) {
  for key, value := range event {
    logp.Info("kv: %s %s\n", key, value)
  }
}
