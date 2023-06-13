import React, { useContext, useEffect, useState } from 'react';
import { Link } from 'react-router-dom';
import videoService from '../../services/video.service';
import { Notify } from 'notiflix/build/notiflix-notify-aio';
import Modal from 'react-modal';
import { AuthContext } from '../../context/AuthContext';
import { FiX } from 'react-icons/fi';
import { TbCrown } from 'react-icons/tb';
import Paginate from '../../components/Paginate';
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

const premiumStyle: React.CSSProperties = {
  position: 'absolute',
  top: '10px',
  right: '10px',
  padding: '5px 10px',
  backgroundColor: '#DAA520',
  color: '#fff',
  fontSize: '12px',
  fontWeight: 'bold',
  zIndex: 999,
};

interface VideoItemData {
  id: number;
  name: string;
  description: string;
  category: string;
  tags: string[];
  thumbnail_url: string;
  IsPremium: boolean;
}

const UserDashboardPage = () => {
  const [videos, setVideos] = useState<VideoItemData[]>([]);
  const [isOpen, setIsOpen] = useState<boolean>(false);
  const { user } = useContext(AuthContext);

  const [currentPage, setCurrentPage] = useState(1);
  const [videosPerPage] = useState(6);

  const indexOfLastVideo = currentPage * videosPerPage;
  const indexOfFirstVideo = indexOfLastVideo - videosPerPage;
  const currentVideos = videos.slice(indexOfFirstVideo, indexOfLastVideo);

  let searchString = '';

  const openModal = () => {
    setIsOpen(true);
  };

  const closeModel = () => {
    setIsOpen(false);
  };

  const alertPrem = () => {
    if (user?.role == 'free') {
      openModal();
    }
  };

  useEffect(() => {
    if (user?.role == 'free') {
      openModal();
    }
    const { request } = videoService.takeVideoRecommended(0);
    request
      .then((res) => {
        const videoData = res.data.map(
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
        setVideos(videoData);
      })
      .catch((err) => {
        console.error(err);
      });
  }, []);

  const handleSearch = (e: any) => {
    e.preventDefault();
    if (!searchString) {
      const { request } = videoService.takeVideoRecommended(0);
      request
        .then((res) => {
          const videoData = res.data.map(
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
          setVideos(videoData);
        })
        .catch((err) => {
          console.error(err);
        });
    } else {
      const { request } = videoService.takeSearchVideo(searchString);
      request
        .then((res) => {
          const videoData = res.data.map(
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
          setVideos(videoData);
        })
        .catch((err) => {
          console.error(err);
        });
    }
  };

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
      <div className="container-fluid p-0 d-flex justify-content-between align-items-center flex-column h-100">
        <div className="container p-4 mx-0">
          <div className="d-flex flex-wrap">
            <form className="col-12" onSubmit={handleSearch}>
              <input
                type="search"
                onChange={(e) => {
                  searchString = e?.target.value;
                }}
                className="form-control ps-8"
                placeholder="Search..."
                aria-label="Search"
              />
            </form>
          </div>
        </div>
        <div className="row row-cols-3 col-8">
          {currentVideos ? (
            currentVideos.map((item, index) => (
              <div className="col my-2" key={index}>
                <Link
                  to={
                    item.IsPremium ? (user?.role != 'free' ? '/VideoPlayer/' + item.id : '') : '/VideoPlayer/' + item.id
                  }
                  onClick={() => alertPrem()}
                  className={`card cardMY justify-content-center ${item.IsPremium ? 'border border-warning' : ''}`}
                >
                  <div className="p-1 myP">
                    <img
                      src={'http://localhost:9000/thumbnails/' + item.thumbnail_url}
                      className="card-img-top myImg"
                      alt="logo kursu"
                    />
                    {item.IsPremium && (
                      <span style={premiumStyle} className="premium-badge">
                        Premium
                      </span>
                    )}
                    <div className="card-body">
                      <div className="card-text">
                        <p className="lead">
                          {item.name}
                          {item.IsPremium && <TbCrown style={{ color: '#DAA520' }} />}
                        </p>
                        <p className="myDes">{`${item.description.substring(0, 35)}...`}</p>
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
        <div className="d-flex justify-content-center p-4">
          <Paginate postsPerPage={6} totalPosts={videos.length} paginate={setCurrentPage} />
        </div>
      </div>
    </>
  );
};

export default UserDashboardPage;
