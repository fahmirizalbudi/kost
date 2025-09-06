"use client"

import Button from "@/app/components/ui/Button"
import styles from "./AppBar.module.scss"
import { asset } from "@/app/lib/asset"
import Image from "next/image"

const AppBar = () => {
  return (
    <nav className={styles.appbar}>
      <span className={styles.margnify}>
        <Image src={asset("magnify.svg")} className={styles.magnify} alt="Search" width={21} height={21} />
      </span>
      <div>
        <Button className={styles.profile}>
          <Image src={asset("profile.svg")} alt="Profile" width={16} height={16} />
        </Button>
      </div>
    </nav>
  )
}

export default AppBar
