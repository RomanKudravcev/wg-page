export default function ShoppingList() {
  return (
    <div class="grid grid-cols-1  bg-white border border-gray-200 rounded-lg shadow w-full">
      <a href="#">
        <h5 class="mb-2 text-2xl font-bold tracking-tight text-gray-900 ">
          Einkaufsliste
        </h5>
      </a>
      <div className="divide-y text-start">
        <div class="flex items-center ml-3 py-2 gap-4">
          <input
            id="offers"
            name="offers"
            type="checkbox"
            className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
          />
          <p class="font-normal text-gray-700">Tomate</p>
        </div>
        <div class="flex items-center ml-3 py-4 gap-4">
          <input
            id="offers"
            name="offers"
            type="checkbox"
            className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
          />
          <p class="font-normal text-gray-700">Apfel</p>
        </div>
        <div class="flex items-center ml-3 py-4 gap-4">
          <input
            id="offers"
            name="offers"
            type="checkbox"
            className="h-4 w-4 rounded border-gray-300 text-indigo-600 focus:ring-indigo-600"
          />
          <p class="font-normal text-gray-700">Apfel</p>
        </div>
        <div></div>
      </div>
      <div className="flex p-1 gap-1">
        <input
          id="email"
          name="email"
          type="email"
          autoComplete="email"
          required
          className="block w-3/4 rounded-md border-0 py-1.5 text-gray-900 shadow-sm ring-1 ring-inset ring-gray-300 placeholder:text-gray-400 focus:ring-2 focus:ring-inset focus:ring-indigo-600 sm:text-sm sm:leading-6"
        />{" "}
        <button
          type="submit"
          className="flex w-1/4 justify-center rounded-md bg-indigo-600 px-3 py-1.5 text-sm font-semibold leading-6 text-white shadow-sm hover:bg-indigo-500 focus-visible:outline focus-visible:outline-2 focus-visible:outline-offset-2 focus-visible:outline-indigo-600"
        >
          Add
        </button>
      </div>
    </div>
  );
}
