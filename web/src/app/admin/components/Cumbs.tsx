import styles from "./Cumbs.module.scss"

type CumbsProps = {
  heading: string
  description: string
}

const Cumbs = (cumbsProps: CumbsProps) => {
  return (
    <>
      <p className={styles.heading}>{cumbsProps.heading}</p>
      <p className={styles.description}>{cumbsProps.description}</p>
    </>
  )
}

export default Cumbs
