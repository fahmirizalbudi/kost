"use client"

import { useState } from "react"
import styles from "./AddUser.module.scss"
import { Modal } from "@/app/components/ui/Modal"
import Button from "@/app/components/ui/Button"
import TextBox from "@/app/components/forms/TextBox"
import Flex from "@/app/components/layout/Flex"
import Label from "@/app/components/forms/Label"
import ComboBox, { Option } from "@/app/components/forms/ComboBox"
import FormAction from "@/app/components/forms/FormAction"
import { findOption } from "@/app/utils/utils"
import { User } from "@/app/types/user"
import Error from "@/app/components/forms/Error"
import { API_URL } from "@/app/constants/api"
import { useRouter } from "next/navigation"

const AddUser = () => {
  const router = useRouter()
  const [isOpen, setIsOpen] = useState<Boolean>(false)
  const [formData, setFormData] = useState<User>({ name: "", email: "", role: "", password: "", phone: "", address: "" })
  const [errors, setErrors] = useState<User>({ name: "", email: "", role: "", password: "", phone: "", address: "" })

  const toggleOpen = () => {
    setFormData({ name: "", email: "", role: "", password: "", phone: "", address: "" })
    setErrors({ name: "", email: "", role: "", password: "", phone: "", address: "" })
    return setIsOpen(!isOpen)
  }

  const roles: Option[] = [
    { value: "admin", label: "Administrator" },
    { value: "tenant", label: "Penyewa" },
  ]

  const handleChange = (e: React.ChangeEvent<HTMLInputElement>) => {
    setFormData({ ...formData, [e.target.name]: e.target.value })
  }

  const handleSubmit = async (e: React.FormEvent<HTMLFormElement>) => {
    e.preventDefault()
    const res = await fetch(API_URL + "/users", {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
        Accept: "application/json",
      },
      body: JSON.stringify(formData),
    })
    const json = await res.json()

    if (res.status === 422) {
      const errs = json.data
      return setErrors(errs)
    }

    setFormData({
      name: "",
      email: "",
      role: "",
      password: "",
      phone: "",
      address: "",
    })

    setIsOpen(false)
    router.refresh()
  }

  return (
    <>
      <Button className={styles.add} onClick={toggleOpen}>
        Tambah Pengguna
      </Button>
      <Modal title="Tambah Pengguna" isOpen={isOpen} onClose={toggleOpen}>
        <form className={styles.form} onSubmit={handleSubmit}>
          <Flex className={styles.group}>
            <Label htmlFor="name">Nama :</Label>
            <TextBox name="name" type="text" value={formData.name} placeholder="Masukkan nama ..." onChange={handleChange} />
            <Error error={errors.name} />
          </Flex>
          <Flex className={styles.group}>
            <Label htmlFor="email">Email :</Label>
            <TextBox name="email" type="email" value={formData.email} placeholder="Masukkan email ..." onChange={handleChange} />
            <Error error={errors.email} />
          </Flex>
          <Flex className={styles.group}>
            <Label htmlFor="role">Hak :</Label>
            <ComboBox
              options={roles}
              placeholder="Pilih hak akses ..."
              value={findOption(roles, formData.role)}
              onChange={(option) => setFormData({ ...formData, role: option })}
            />
            <Error error={errors.role} />
          </Flex>
          <Flex className={styles.group}>
            <Label htmlFor="password">Password :</Label>
            <TextBox name="password" type="text" value={formData.password} placeholder="Masukkan password ..." onChange={handleChange} />
            <Error error={errors.password} />
          </Flex>
          <Flex className={styles.group}>
            <Label htmlFor="phone">Seluler :</Label>
            <TextBox name="phone" type="text" value={formData.phone} placeholder="Masukkan seluler ..." onChange={handleChange} />
            <Error error={errors.phone} />
          </Flex>
          <Flex className={styles.group}>
            <Label htmlFor="address">Alamat :</Label>
            <TextBox name="address" type="text" value={formData.address} placeholder="Masukkan alamat ..." onChange={handleChange} />
            <Error error={errors.address} />
          </Flex>
          <FormAction onCancel={toggleOpen} />
        </form>
      </Modal>
    </>
  )
}

export default AddUser
