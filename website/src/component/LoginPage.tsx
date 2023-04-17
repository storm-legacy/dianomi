import React, { useState } from "react";
function LoginPage() {
  const [username, setUsername] = useState("");
  const [password, setPassword] = useState("");

  function handleSubmit(event: any) {
    event.preventDefault();
    if (username === "admin" && password === "admin") {
      console.log("Użytkownik zalogowany");
    } else {
      console.log("Błędne dane logowania");
    }
  }

  return (
    <div className="Login">
      <h3>Logowanie</h3>
    <form onSubmit={handleSubmit}>
      <label>
        Nazwa użytkownika:<br/>
        <input
          type="text"
          value={username}
          onChange={(event) => setUsername(event.target.value)}
        />
      </label>
      <br />
      <label>
        Hasło:<br/>
        <input
          type="password"
          value={password}
          onChange={(event) => setPassword(event.target.value)}
        />
      </label>
      <br />
      <button type="submit">Zaloguj</button>
      

    </form>
    </div>
  );
}

export default LoginPage;
