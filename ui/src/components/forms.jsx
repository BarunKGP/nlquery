"use client";

import { z } from "zod";
import React from "react";

const formSchema = z.object({
  text: z.string().max(1000),
});

function Forms() {
  return <div>forms</div>;
}

export default Forms;
