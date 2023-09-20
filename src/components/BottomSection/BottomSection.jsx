import React from 'react';
import { useState, useEffect } from 'react';
import TwitchPage from './TwitchPage';
import PresentationPage from './PresentationPage';


const BottomSection = () => {
	const [twitchData, setTwitchData] = useState(undefined);

	useEffect(() => {
		if (twitchData !== undefined) return;
		setTwitchData(undefined);
		fetch("https://randomuser.me/api/").then( async (response) => {
			if (response.status < 200 || response.status >= 400) return;
			let body = await response.json();
			setTwitchData(body.results[0]);
		})
	}, [twitchData]);

	return (
		<>
			<PresentationPage/>
			{/* <HobbiesPage/> */}
			{/* <TwitchPage/> */}
			{/* <YoutubePage/> */}
		</>
	);
}

export default BottomSection;