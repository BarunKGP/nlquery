import axios from "axios";
import { getSession } from "next-auth/react";
import { CustomSession } from ".";

const createAuthenticatedClient = async () => {
  const session = ((await getSession()) as CustomSession) || null;
  const token = session?.user?.backendToken;

  return axios.create({
    baseURL: process.env.BACKEND_SERVER,
    withCredentials: true,
    // headers: { Authorization: token ? `Bearer ${token}` : "" },
  });
};

export default createAuthenticatedClient;
