
import React, { FC } from 'react';
import { Link as RouterLink } from 'react-router-dom';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
 Link,
} from '@backstage/core';
import Button from '@material-ui/core/Button';
import ComponanceTable from '../Table';
import AccountCircleIcon from '@material-ui/icons/AccountCircle';

const PageResult: FC<{}> = () => {

 return (
   <Page theme={pageTheme.app}>
     <Header
       title={'ระบบการจัดการรถพยาบาล'}
       subtitle="ระบบหลัก: ระบบรถพยาบาล"
     >
      <td align='center'>
        <AccountCircleIcon style={{ fontSize: 50}}>AccountCircleIcon</AccountCircleIcon>
        <br></br>
        เจ้าหน้าที่จัดการรถพยาบาล
        <br></br>
        tsaiprom.nams@gmail.com
        <br></br>
        <Link component={RouterLink} to="/">logout
         </Link>
      </td>
     </Header>
     <Content>
      <ContentHeader title="สถานะรถพยาบาลที่ถูกบันทึก">
         <Link component={RouterLink} to="/user">
           <Button variant="contained" color="primary">
             บันทึกสถานะรถ
           </Button>
         </Link>
       </ContentHeader>
       <ComponanceTable></ComponanceTable>
     </Content>
   </Page>
 );
}

export default PageResult;