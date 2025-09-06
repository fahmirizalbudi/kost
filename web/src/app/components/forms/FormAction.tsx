import Button from "../ui/Button"
import styles from "./FormAction.module.scss"

type FormActionProps = {
  onCancel: () => void
}

const FormAction = ({ onCancel }: FormActionProps) => {
  return (
    <footer className={styles.actions}>
      <Button className={styles.submit} type="submit">
        Submit
      </Button>
      <Button className={styles.cancel} type="button" onClick={onCancel}>
        Cancel
      </Button>
    </footer>
  )
}

export default FormAction
