import React from 'react';
import { useNavigate, useSearchParams, Link } from 'react-router-dom';
import developmentService from '../../services/development.service';
import { User } from '../../types/user.type';
import { Notify } from 'notiflix';

const PaymentPage = () => {
  const navigate = useNavigate();
  const [searchParams, setSearchParams] = useSearchParams();
  const status = searchParams.get('status');

  if (status === 'success') {
    const user: User = JSON.parse(String(localStorage.getItem("user")))
    const { request } = developmentService.GivePackageSelf(user.email)
    request.then(() => {
      Notify.info("You can know access premium materials");
    }).catch(() => {
      Notify.failure("Problem occured while giving user the package");
    })

    return (
      <>
        <div className="container d-flex justify-content-center align-items-center h-75">
          <div className="border text-white text-center border-secondary bg-success rounded-2 p-4 col-4 shadow">
            <h2>Thank you for your purchase!</h2>
            <div className="p-4">
              <span>You can now move to application and enjoy premium content.</span>
            </div>
            <Link to="/">
              <button className="btn btn-danger shadow">Return to application</button>
            </Link>
          </div>
        </div>
      </>
    );
  }
  else
    return (
      <>
        <div className="container d-flex justify-content-center align-items-center h-75">
          <div className="border text-center border-secondary bg-warning rounded-2 p-4 col-4 shadow">
            <h2>Payment was not finished</h2>
            <div className="p-4">
              <span>
                Payment operation was canceled or another error occured. In case it was not intentional - please try
                again.
              </span>
            </div>
            <div>
              <Link to="/">
                <button className="btn btn-danger shadow">Return to application</button>
              </Link>
            </div>
          </div>
        </div>
      </>
    );
};

export default PaymentPage;
