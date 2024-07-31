"use client";
import React, { useState } from "react";
import { Button } from "@/components/ui/button";
import Image from "next/image";

const UploadButton = () => {
  const [selectedFile, setSelectedFile] = useState(null);

  // Function to handle file selection
  const handleFileChange = (event) => {
    setSelectedFile(event.target.files[0]);
  };

  // Function to handle file upload
  const handleUpload = async () => {
    if (!selectedFile) {
      alert("Please select a file to upload.");
      return;
    }

    // Assuming you have a backend endpoint for file upload
    const formData = new FormData();
    formData.append("file", selectedFile);

    await fetch("https://localhost:8000/api/v1/prompt", {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json(); // Assuming server returns JSON response
      })
      .then((data) => {
        console.log("File uploaded successfully:", data);
        // Handle further processing or update UI as needed
      })
      .catch((error) => {
        console.error("Error uploading file:", error);
        // Handle errors or display appropriate message to the user
      });
  };

  return (
    <div>
      <button
        type="submit"
        className="rounded-md px-3 py-2 text-sm font-medium text-primary-background hover:bg-secondary/90 focus-visible:outline-none focus-visible:ring-1 focus-visible:ring-ring"
        onChange={handleFileChange}
      >
        <UploadIcon />
      </button>
    </div>
  );
};

function UploadIcon(props) {
  return (
    <svg
      {...props}
      xmlns="http://www.w3.org/2000/svg"
      viewBox="0 0 24 24"
      height={24}
      width={24}
    >
      <path
        d="M9.707 7.707 11 6.414V16a1 1 0 0 0 2 0V6.414l1.293 1.293a1 1 0 0 0 1.414-1.414l-3-3a1 1 0 0 0-1.416 0l-3 3a1 1 0 0 0 1.416 1.414z"
        style={{ fill: "#ff8e31" }}
      />
      <path
        d="M17 19H7a1 1 0 0 0 0 2h10a1 1 0 0 0 0-2z"
        style={{ fill: "#ece4b7" }}
      />
    </svg>
  );
}

export default UploadButton;
