import React, { useState } from 'react';

function RegisterPage() {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');
  const [email, setEmail] = useState('');

  function handleSubmit(event: any) {
    event.preventDefault();
    console.log('Zarejestrowany użytkownik: ', username, email, password);
  }

  return (
    <div className="text-center float-start" style={{ marginLeft: 50 }}>
      <h3>Rejestracja</h3>
      <form onSubmit={handleSubmit}>
        <label>
          <p className="h5">Nazwa użytkownika:</p>
          <input
            type="text"
            className="form-control"
            placeholder="Username"
            aria-label="Username"
            aria-describedby="basic-addon1"
            value={username}
            onChange={(event) => setUsername(event.target.value)}
          />
        </label>
        <br />
        <label>
          <p className="h5">Adres email:</p>
          <input
            type="email"
            className="form-control"
            aria-describedby="emailHelp"
            placeholder="Enter Email"
            value={email}
            onChange={(event) => setEmail(event.target.value)}
          />
        </label>
        <br />
        <label>
          <p className="h5">Hasło:</p>
          <input
            type="password"
            className="form-control"
            placeholder="Password"
            aria-label="Password"
            aria-describedby="basic-addon1"
            value={password}
            onChange={(event) => setPassword(event.target.value)}
          />
        </label>
        <br />
        <br />
        <button type="submit" className="btn btn-primary">
          Zarejestruj
        </button>
      </form>
    </div>
  );
}

export default RegisterPage;
