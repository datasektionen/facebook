import styles from '../styles/style.module.css'

export default function RootLayout({
  children,
}: {
  children: React.ReactNode
}) {
  return (
    <>
      <main className={styles.global}>{children}</main>
    </>
  )
}
