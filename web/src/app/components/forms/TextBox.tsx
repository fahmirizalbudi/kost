import styles from "./TextBox.module.scss"

type TextBoxProps = {
  type: "text" | "email" | "number" | "hidden"
  placeholder: string
  onChange?: (e: React.ChangeEvent<HTMLInputElement>) => void
  name?: string
  value?: number | string
}

const TextBox = ({ type, placeholder, onChange, name, value }: TextBoxProps) => (
  <input className={styles.textbox} type={type} placeholder={placeholder} onChange={onChange} name={name} value={value} id={name} autoComplete="off" />
)

export default TextBox
