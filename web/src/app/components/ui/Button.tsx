type ButtonProps = {
    className?: string,
    onClick?: () => void,
    children: React.ReactNode,
    type?: "button" | "submit"
}

const Button = (buttonProps: ButtonProps) => {
  return (
    <button className={buttonProps.className} onClick={buttonProps.onClick} type={buttonProps.type}>{buttonProps.children}</button>
  )
}

export default Button