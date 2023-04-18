import React, { useState } from 'react';
import { Link } from 'react-router-dom';

const LoginPage = () => {
  const [username, setUsername] = useState('');
  const [password, setPassword] = useState('');

  const handleSubmit = (event: any) => {
    event.preventDefault();
  };

  return (
    <div className="text-center float-start">
      <h3>Logowanie</h3>
      <form onSubmit={handleSubmit}>
        <label>
          <p className="h5">Nazwa użytkownika:</p>
          <br />
          <input
            type="text"
            className="form-control"
            placeholder="UserName"
            aria-label="UserName"
            aria-describedby="basic-addon1"
            value={username}
            onChange={(event) => setUsername(event.target.value)}
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
        <button type="submit" className="btn btn-primary" onSubmit={handleSubmit}>
          Zaloguj
        </button>
      </form>
      <Link to={'/Register'}>Rejestracja</Link>
    </div>
  );
};

export default LoginPage;
