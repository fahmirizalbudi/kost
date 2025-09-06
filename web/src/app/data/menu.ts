export type menuNavigationBarProps = {
  menu: string
  linkTo: string
}

export type menuSideNavigationProps = {
  text: string
  linkTo?: string
  icon?: string
  type: "header" | "link"
}

export const MENU_SIDE_HEADER = "header"
export const MENU_SIDE_LINK = "link"

export const menuNavigationBar: menuNavigationBarProps[] = [
  {
    menu: "Beranda",
    linkTo: "/",
  },
  {
    menu: "Sewa",
    linkTo: "/rent",
  },
  {
    menu: "Kontak",
    linkTo: "/contact",
  },
  {
    menu: "Telusuri",
    linkTo: "/search",
  },
]

export const menuSideNavigation: menuSideNavigationProps[] = [
  {
    text: "MAIN MENU",
    type: MENU_SIDE_HEADER
  },
  {
    text: "Beranda",
    linkTo: "/admin",
    icon: "beranda.svg",
    type: MENU_SIDE_LINK
  },
  {
    text: "Pengguna",
    linkTo: "/admin/users",
    icon: "pengguna.svg",
    type: MENU_SIDE_LINK
  },
  {
    text: "Kost",
    linkTo: "/admin/dormitories",
    icon: "kost.svg",
    type: MENU_SIDE_LINK
  },
  {
    text: "Kamar Kost",
    linkTo: "/admin/rooms",
    icon: "kamar.svg",
    type: MENU_SIDE_LINK
  },
  {
    text: "OTHERS",
    type: MENU_SIDE_HEADER
  },
  {
    text: "Penyewaan",
    linkTo: "/admin/rentals",
    icon: "penyewaan.svg",
    type: MENU_SIDE_LINK
  },
  {
    text: "Transaksi",
    linkTo: "/admin/transactions",
    icon: "transaksi.svg",
    type: MENU_SIDE_LINK
  },
]
