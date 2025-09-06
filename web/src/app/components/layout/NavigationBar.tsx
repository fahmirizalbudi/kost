import { menuNavigationBarProps } from "@/app/data/menu"
import { asset } from "@/app/lib/asset"
import Image from "next/image"
import Button from "../ui/Button"
import styles from "./NavigationBar.module.scss"

type NavigationBarProps = {
    menu: menuNavigationBarProps[]
}

const NavigationBar = ({ menu }: NavigationBarProps) => {
  return (
    <nav className={styles.nav}>
        <Image src={asset("logo.png")} alt="Logo" width={35} height={35} />
        <ul className={styles.menu}>
            {menu.map((item, i) => (
                <li key={i} className={styles.i}><a href={item.linkTo} className={styles.anchor}>{item.menu}</a></li>
            ))}
        </ul>
        <Button className={styles.profile}>
            <Image src={asset("profile.svg")} alt="Profile" width={16} height={16} />
        </Button>
    </nav>
  )
}

export default NavigationBar