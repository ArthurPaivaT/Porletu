import "./App.css";
import NavBar from "./Components/NavBar/NavBar";
import SignUpPage from "./Pages/SignUp/signup";

function App() {
  return (
    <>
      <NavBar />
      <div className="appArea">
        <SignUpPage className="UserCard" />
      </div>
    </>
  );
}

export default App;
