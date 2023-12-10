export default function Navbar() {
  return (
    <header className="flex items-center text-[#8F9EC4] bg-indigo-100 p-2 pr-6  h-[10vh] ">
      {/*Logo*/}
      <div className="flex flex-initial items-center h-full w-3/12 ">
        <a href="/dashboard" className=" h-full">
          <img src="/logo.png" className="h-full" />
        </a>
        <a
          href="/dashboard"
          className="font-sans text-5xl font-semibold pl-3 hover:text-[#687BA8]"
        >
          WG
        </a>
      </div>
      {/*Links*/}
      <div className="flex w-6/12 gap-6 text-xl">
        <a href="/todo" className=" hover:text-[#687BA8]">
          <span>TODO</span>
        </a>
        <a href="/help" className=" hover:text-[#687BA8]">
          <span>HELP</span>
        </a>
      </div>
      <div className="flex justify-end w-3/12">
        <a
          href="#"
          className="rounded-lg px-3 py-1 text-base font-semibold text-gray-900 bg-indigo-300 hover:bg-indigo-400"
        >
          Log in
        </a>
      </div>
    </header>
  );
}
