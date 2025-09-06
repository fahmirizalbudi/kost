import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from "../components/Table"
import Cumbs from "../components/Cumbs"
import SafeView from "../components/SafeView"
import styles from "./page.module.scss"
import Button from "@/app/components/ui/Button"
import Image from "next/image"
import { asset } from "@/app/lib/asset"
import { API_URL } from "@/app/constants/api"
import Break from "../components/Break"

const fetchUsers = async () => {
  const res = await fetch(API_URL + "/users")
  const json = await res.json()
  return json.data
}

const Users = async () => {
  const users = await fetchUsers()

  return (
    <SafeView>
      <Cumbs heading="Pengguna" description="Pusat data pengguna untuk melihat, menambah, atau mengelola akun." />
      <Break height={30} />
      <Table>
        <TableHeader>
          <TableRow>
            <TableHead>NO</TableHead>
            <TableHead>NAMA</TableHead>
            <TableHead>EMAIL</TableHead>
            <TableHead>SELULER</TableHead>
            <TableHead>ALAMAT</TableHead>
            <TableHead>AKSI</TableHead>
          </TableRow>
        </TableHeader>
        <TableBody>
          {users.map((user: any, i: number) => (
            <TableRow key={user.id}>
              <TableCell>{i + 1}</TableCell>
              <TableCell>{user.name}</TableCell>
              <TableCell>{user.email}</TableCell>
              <TableCell>{user.phone}</TableCell>
              <TableCell>{user.address}</TableCell>
              <TableCell className={styles.flex}>
                <Button className={`${styles.action}`}>
                  <Image src={asset("edit.svg")} alt="Edit" width={18} height={18} />
                </Button>
                <Button className={`${styles.action}`}>
                  <Image src={asset("hapus.svg")} alt="Delete" width={18} height={18} />
                </Button>
              </TableCell>
            </TableRow>
          ))}
        </TableBody>
      </Table>
    </SafeView>
  )
}

export default Users
