"use client";
import { useSession } from "next-auth/react";

function Page() {
  const session = useSession();
  console.log(session);

  session.data ? (
    <div className="text-gray-800">
      Logged in session data: {session.data.user}
    </div>
  ) : (
    <div className="text-gray-900">Not logged in</div>
  );
}

export default Page;
