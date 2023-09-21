import React from 'react';
import { useState, useEffect } from 'react';
// import TwitchPage from './TwitchPage';
import PresentationPage from './PresentationPage';


const BottomSection = () => {
	const [schoolData, setSchoolData] = useState(undefined);

	useEffect(() => {
		if (schoolData !== undefined) return;
		setSchoolData(undefined);
		fetch("https://randomuser.me/api/").then( async (response) => {
			if (response.status < 200 || response.status >= 400) return;
			let body = await response.json();
			setSchoolData(body.results[0]);
		})
	}, [schoolData]);

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