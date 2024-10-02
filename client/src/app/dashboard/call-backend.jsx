// This is a toy function meant to check if we can call the BACKEND_SERVER
// from the frontend. Mainly used to debug connection issues, CORS, etc.
//! Once you are confident everything works, delete this
import React from "react";
import { auth } from "@/auth";

async function CallBackend() {
  const session = await auth();
  // const data = await axios.post(`${process.env.BACKEND_SERVER}/auth/signin`)

  const res = await fetch(`${process.env.BACKEND_SERVER}/auth/signin`, {
    method: "POST",
    body: JSON.stringify(session.user),
    credentials: "include",
  });
  const data = await res.json();

  return (
    <div className="p-2 mt-5 tracking-tighter text-center bg-yellow-400 border-4 border-orange-600 bg-opacity-90 w-fit text-background rounded-xl">
      <h1>Calling backend</h1>
      <p className="text-xs text-primary w-max-[200px]">
        {JSON.stringify(data, null, 4)}
      </p>
    </div>
  );
}

export default CallBackend;
