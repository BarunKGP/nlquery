"use client";

import { signIn, signOut, useSession } from "next-auth/react";

export default function AuthButton() {
  const { session } = useSession();

  if (session) {
    return (
      <>
        {session.user.email} <br />
        <button onClick={() => signOut()}>Sign out</button>
      </>
    );
  }
  return (
    <>
      Not signed in <br />
      <button onClick={() => signIn()}>Sign in</button>
    </>
  );
}
