import styles from "./page.module.css";
import Link from "next/link";

export default async function Home() {
  return (
    <div className={styles.page}>
      <Link href="/users">Users</Link>
    </div>
  );
}
