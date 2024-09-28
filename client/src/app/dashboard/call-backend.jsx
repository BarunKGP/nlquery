// This is a toy function meant to check if we can call the BACKEND_SERVER
// from the frontend. Mainly used to debug connection issues, CORS, etc.
//! Once you are confident everything works, delete this
import React from "react";

async function CallBackend() {
  const res = await fetch(process.env.BACKEND_SERVER);
  const data = await res.json();

  return (
    <div className="p-2 mt-5 tracking-tighter bg-yellow-400 border-4 border-orange-600 bg-opacity-90 w-fit text-background rounded-xl">
      <h1>Calling backend</h1>
      <p>{data.message}</p>
    </div>
  );
}

export default CallBackend;
