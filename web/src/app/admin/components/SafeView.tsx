import styles from "./SafeView.module.scss"

type SafeViewProps = {
  children?: React.ReactNode
}

const SafeView = ({ children }: SafeViewProps) => {
  return <main className={styles.view}>{children}</main>
}

export default SafeView
