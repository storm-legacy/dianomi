import React, { useContext, useState } from 'react';
import { AuthContext } from '../../context/AuthContext';
import ProfileService, { emailData } from '../../services/profile.service';
import { useNavigate } from 'react-router-dom';
import { Notify, Report } from 'notiflix';

export const ReportPage = () => {
  const navigate = useNavigate();
  const [title, setTitle] = useState<string>('');
  const [dsescription, setDescription] = useState<string>('');
  const [disabled, setDisabled] = useState(false);
  const { user } = useContext(AuthContext);
  const handleSubmit = async (event: any) => {
    event.preventDefault();

    const data: emailData = {
      ErrorTitle: title,
      ErrorDescription: dsescription,
      ReportedBy: String(user?.email),
    };

    setDisabled(true);
    const { request } = ProfileService.PostRaport(data);
    request.then(() => {
      Report.success("Report sent", "Thank you for your feedback.", "Ok", () => {
        navigate("/ProfilePage")
      });
    }).catch(() => {
      Report.failure("Sending error", "Report could not be sent, please try again later.", "Ok");
    })
  };

  return (
    <>
      <div className="position-absolute top-50 start-50 translate-middle text-center float-start shadow-lg p-3 mb-5 bg-white rounded">
        <h3>Error Reports</h3>
        <p></p>
        <form className="row">
          <label>
            <p>Title</p>
            <input className="form-control" type="text" value={title} onChange={(e) => setTitle(e.target?.value)} />
          </label>
          <label>
            <p>Description</p>
            <textarea
              className="form-control"
              style={{ height: '15dvh' }}
              id="exampleFormControlTextarea1"
              value={dsescription}
              onChange={(e) => setDescription(e.target?.value)}
            ></textarea>
          </label>
          <p></p>
          <button type="submit" disabled={disabled} onClick={handleSubmit} className="btn btn-danger">
            Report an Error
          </button>
        </form>
      </div>
    </>
  );
};
