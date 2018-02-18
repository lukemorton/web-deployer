package publish

import(
  "fmt"
  "os/exec"
)

func isExecutableInstalled(name string) bool {
  _, err := exec.LookPath(name)
  return err == nil
}

func runExecutable(executable string, args ...string) error {
  out, err := exec.Command(executable, args...).Output()
  if err != nil {
    return fmt.Errorf("%s", out)
  }

  return nil
}
