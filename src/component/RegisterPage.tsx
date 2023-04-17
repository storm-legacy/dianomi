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
    <div className="Register">
      <h3>Rejestracja</h3>
      <form onSubmit={handleSubmit}>
        <label>
          Nazwa użytkownika: <br />
          <input type="text" value={username} onChange={(event) => setUsername(event.target.value)} />
        </label>
        <br />
        <label>
          Adres email:
          <br />
          <input type="email" value={email} onChange={(event) => setEmail(event.target.value)} />
        </label>
        <br />
        <label>
          Hasło:
          <br />
          <input type="password" value={password} onChange={(event) => setPassword(event.target.value)} />
        </label>
        <br />
        <button type="submit">Zarejestruj</button>
      </form>
    </div>
  );
}

export default RegisterPage;
