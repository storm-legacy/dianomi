import React, { useContext, useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import videoService from '../../services/video.service';
import { Notify } from 'notiflix/build/notiflix-notify-aio';
import Modal from 'react-modal';
import { AuthContext } from '../../context/AuthContext';
import { FiX } from 'react-icons/fi';
import { TbCrown } from 'react-icons/tb';
const customStyles = {
  overlay: {
    background: 'none',
  },
  content: {
    left: '75%',
    height: '35dvh',
    width: '23dvw',
  },
};

const UserDashboardPage = () => {
  interface VideoItemData {
    id: number;
    name: string;
    description: string;
    category: string;
    tags: string[];
    thumbnail_url: string;
    IsPremium: boolean;
  }
  const [divItem, setDivItem] = useState<VideoItemData[]>([]);
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const { user } = useContext(AuthContext);
  function openModal() {
    setIsOpen(true);
  }
  function closeModel() {
    setIsOpen(false);
  }

  useEffect(() => {
    if (user?.role == 'free') {
      openModal();
    }
    const { request } = videoService.takeVideoRecommended();
    request
      .then((res) => {
        console.log(res);
        const Videodata = res.data.map(
          (Videodata: {
            id: number;
            name: string;
            description: string;
            category: string;
            tags: string[];
            thumbnail_url: string;
            IsPremium: boolean;
          }) => {
            return {
              id: Videodata.id,
              name: Videodata.name,
              description: Videodata.description,
              category: Videodata.category,
              tags: Videodata.tags,
              thumbnail_url: Videodata.thumbnail_url,
              IsPremium: Videodata.IsPremium,
            };
          },
        );
        setDivItem(Videodata);
      })
      .catch((err) => {
        console.error(err);
      });
  }, []);

  return (
    <>
      <Modal style={customStyles} isOpen={isOpen} onRequestClose={closeModel}>
        <div className="text-center ">
          <FiX onClick={() => closeModel()} style={{ float: 'right' }}></FiX>
          <h3>Hello friend</h3>
          <p>You can purchase a premium package giving you access to a larger video library on our website.</p>{' '}
          <h6>4.99PLN/month</h6>
          <a href="/ProfilePage">
            <button className="btn btn-danger">Buy Now</button>
          </a>
        </div>
      </Modal>
      <div className="container m-0 p-4">
        <div className="row row-cols-3 ">
          {divItem ? (
            divItem.map((item, index) => (
              <div className="col" key={index}>
                <Link to={'/VideoPlayer/' + item.id} className="card cardMY justify-content-center">
                  <div className="p-2 myP">
                    <img
                      src={'http://localhost:9000/thumbnails/' + item.thumbnail_url}
                      className="card-img-top myImg"
                      alt="logo kursu"
                    />
                    <div className="card-body">
                      <div className="card-text">
                        <p className="lead">
                          {item.name}
                          {item.IsPremium && <TbCrown />}
                        </p>
                        <p className="myDes">{`${item.description.substring(0, 124)}...`}</p>
                      </div>
                    </div>
                  </div>
                </Link>
              </div>
            ))
          ) : (
            <span>No videos to show</span>
          )}
        </div>
      </div>
    </>
  );
};

export default UserDashboardPage;
