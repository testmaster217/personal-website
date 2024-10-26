"use client"

import { usePathname } from "next/navigation";
import { useEffect, useRef } from "react"

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

    return <dialog ref={pageDialog}>
        <button autoFocus type="button" onClick={() => pageDialog.current && pageDialog.current.close()}>X</button>
        {children}
    </dialog>
}