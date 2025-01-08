"use client"

import Link from "next/link";

export default function NotFound() {
    return (
        <div>
            <p>404 not found</p>
            <Link href="/top">TOPへ</Link>
        </div>
    );
}