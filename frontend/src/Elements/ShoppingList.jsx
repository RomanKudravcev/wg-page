import { useEffect, useState } from "react";

export default function ShoppingList() {
  const [data, setData] = useState(null);
  const [item, setItem] = useState("");

  useEffect(() => {
    fetch("http://localhost:8080/items")
      .then((response) => response.json())
      .then((actualData) => setData(actualData.items));

    console.log(data);
  }, []);

  const handleSubmit = async (event) => {
    event.preventDefault();

    try {
      const response = await fetch("http://localhost:8080/items", {
        method: "POST",
        headers: {
          "Content-Type": "application/json",
        },
        body: JSON.stringify({ item }),
      });

      if (response.ok) {
        console.log("Item added successfully");
        setData([
          ...data,
          {
            item: item,
            bought: false,
          },
        ]);
        setItem("");
        // Optionally reset the form or handle success
      } else {
        console.error("Error adding item");
        // Optionally handle error
      }
    } catch (error) {
      console.error("Error:", error);
    }
  };

  const items = data?.map((item) => {
    return (
      <div className="flex items-center ml-3 py-2 gap-4">
        <input
          id="item"
          name="item"
          type="checkbox"
          className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
        />
        <p className="font-normal text-gray-700">{item.item}</p>
      </div>
    );
  });

  return (
    <div className="grid grid-cols-1  bg-white border border-gray-200 rounded-lg shadow w-full h-2/4">
      <a href="#">
        <h5 className="mb-2 text-2xl font-bold tracking-tight text-gray-900 ">
          Einkaufsliste
        </h5>
      </a>
      <div className="divide-y text-start overflow-y-auto h-full">
        {items}
        <div></div>
      </div>
      <form onSubmit={handleSubmit}>
        <div className="flex p-1 gap-1">
          <input
            name="item"
            type="text"
            required
            value={item}
            onChange={(e) => setItem(e.target.value)}
            className="block w-3/4 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
          />
          <button
            type="submit"
            className="flex w-1/4 justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
          >
            Add
          </button>
        </div>
      </form>
    </div>
  );
}
