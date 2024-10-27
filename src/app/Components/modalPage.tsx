"use client"

import Link from "next/link";

import { usePathname } from "next/navigation";
import { useEffect, useRef } from "react"

import styles from "./modalPage.module.css";

export default function ModalPage({
    children
}: Readonly<{
    children: React.ReactNode
}>) {
    const pageDialog = useRef<HTMLDialogElement>(null);
    const pathname = usePathname();

    useEffect(() => {
        if (pathname.match(/\/./) && pageDialog.current && !pageDialog.current.open)
            pageDialog.current && pageDialog.current.showModal();
    }, [pathname]);

    const returnLink = pathname.substring(0, pathname.lastIndexOf("/")) === "" ? "/" : pathname.substring(0, pathname.lastIndexOf("/"));

    return <dialog ref={pageDialog} className={styles.ModalPage}>
        <Link autoFocus href={returnLink} className={styles.CloseBtn} onClick={() => pageDialog.current && pageDialog.current.close()}>X</Link>
        {children}
    </dialog>
}