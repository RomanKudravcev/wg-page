import React, { useState, useEffect } from "react";
import Navbar from "../Elements/Navbar";
import ShoppingList from "../Elements/ShoppingList";

export default function Dashboard() {
  const [message, setMessage] = useState("");
  const [error, setError] = useState(null);

  useEffect(() => {
    fetchData();
  }, []);

  function fetchData() {
    fetch("http://localhost:8080/api/data")
      .then((response) => {
        if (!response.ok) {
          throw new Error("Network response was not ok");
        }
        return response.json();
      })
      .then((data) => {
        // Assuming the message is in data.message
        setMessage(data.message);
      })
      .catch((error) => {
        console.error("Error fetching data: ", error);
        setError(error);
      });
  }

  return (
    <div>
      <Navbar />
      <div className="flex gap-5 p-5">
        <div className="flex w-3/12">
          <ShoppingList />
        </div>
      </div>
    </div>
  );
}
