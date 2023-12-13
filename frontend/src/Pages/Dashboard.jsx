import Navbar from "../Elements/Navbar";
import ShoppingList from "../Elements/ShoppingList";

export default function Dashboard() {
  return (
    <div>
      <Navbar />
      <div className="flex gap-5 p-5">
        <div className="flex w-3/12 h-5/6">
          <ShoppingList />
        </div>
      </div>
    </div>
  );
}
