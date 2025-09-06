type LabelProps = {
  htmlFor: string
  children: React.ReactNode
}

const Label = ({ htmlFor, children }: LabelProps) => <label style={{ position: "relative", left: 1 }} htmlFor={htmlFor}>{children}</label>

export default Label
