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
