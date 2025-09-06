import styles from "./Error.module.scss"

type ErrorProps = {
  error?: string
}

const Error = ({ error }: ErrorProps) => {
  if (!error) return null
  return (
    <small className={styles.error}>
      <span className={styles.danger}>⚠︎</span> {error}
    </small>
  )
}

export default Error
