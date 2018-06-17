import React from 'react';
import FullLogo from '../../common/logo/full';
import IconHeader from '../../common/iconHeader';
import { length } from '../../../style';

const style = {
  container: {
    margin: length.medium,
    display: 'flex',
    flexDirection: 'column',
    textAlign: 'left'
  },
  logo: {
    marginTop: length.large,
    marginLeft: '5%',
    width: '80%'
  }
}

const About = props => (
  <div style={style.container}>
    <div style={style.logo}>
      <FullLogo />
    </div>
    <p>Diplomacy is a game of international intrigue, trust and treachery.</p>
    <p>At the start of 1901 you are given commmand of the foreign policy of one of the 7 major european powers: Austria, England, France, Germany, Italy, Russia and Turkey.</p>
    <p>Your goal is to take controll of half the map of Europe, a task that cannot be achieved only by military might but will require diplomatic skill and subtelty.</p>
    <p>Luck will play no part, only skill can keep your nation alive, only skill can lead you to hegemony over Europe.</p>
    <p>Always remember what the ancient Athenians knew and reminded their enemies of:</p>
    <i><p>The strong do what they can while the weak suffer what they must.</p></i>
    <IconHeader text='Rules' />
  </div>
);

export default About;
