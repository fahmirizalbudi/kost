import { Option } from "../components/forms/ComboBox";

export function findOption(options: Option[], value: string | undefined | null): Option | null {
  return options.find((option) => option.value === value) || null
}
