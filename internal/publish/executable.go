package publish

import(
  "fmt"
  "os/exec"
)

func ensureExecutableInstalled(name string) error {
  _, err := exec.LookPath(name)

  if err !=nil {
    return fmt.Errorf("Could not find %s executable, is it installed?", name)
  }

  return nil
}

func runExecutable(executable string, args ...string) error {
  out, err := exec.Command(executable, args...).CombinedOutput()
  if err != nil {
    return fmt.Errorf("%s", out)
  }

  return nil
}
