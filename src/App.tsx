import React from 'react';
import LoginPage from './component/LoginPage';
import RegisterPage from './component/RegisterPage';
function App() {
  return (
    <div className="justify-content-center">
      <div className="justify-content-center text-center shadow-lg p-3 mb-5 bg-body rounded d-inline-flex p-2 bd-highlight">
        <LoginPage />
        <RegisterPage />
      </div>
    </div>
  );
}

export default App;
