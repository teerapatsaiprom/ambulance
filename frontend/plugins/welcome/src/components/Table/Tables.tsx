import React, { useState, useEffect } from 'react';
import { makeStyles } from '@material-ui/core/styles';
import Table from '@material-ui/core/Table';
import TableBody from '@material-ui/core/TableBody';
import TableCell from '@material-ui/core/TableCell';
import TableContainer from '@material-ui/core/TableContainer';
import TableHead from '@material-ui/core/TableHead';
import TableRow from '@material-ui/core/TableRow';
import Paper from '@material-ui/core/Paper';
import Button from '@material-ui/core/Button';
import { DefaultApi } from '../../api/apis';
import { EntPredicament } from '../../api/models/EntPredicament';
import moment from 'moment';
 
const useStyles = makeStyles({
 table: {
   minWidth: 650,
 },
});
 
export default function ComponentsTable() {
 const classes = useStyles();
 const api = new DefaultApi();

 
 const [predicament, setPredicaments] = useState<EntPredicament[]>([]);
 const [loading, setLoading] = useState(true);
 
 useEffect(() => {
  const getPredicaments = async () => {
    const res = await api.listPredicament({ limit: 10, offset: 0 });
    setLoading(false);
    setPredicaments(res);
    console.log(res);
  };
  getPredicaments();

}, [loading]);
 
 const deletePredicaments = async (id: number) => {
   const res = await api.deletePredicament({ id: id });
   setLoading(true);
 };
 
 return (
   <TableContainer component={Paper}>
     <Table className={classes.table} aria-label="simple table">
       <TableHead>
         <TableRow>
            <TableCell align="center">บันทึกลำดับ</TableCell>
            <TableCell align="center">Car No.xx (หมายเลขประจำรถ)</TableCell>
            <TableCell align="center">ผู้รับผิดชอบประจำรถ</TableCell>
            <TableCell align="center">สถานะรถพยาบาล</TableCell>
            <TableCell align="center">เจ้าหน้าที่จัดการรถที่บันทึก</TableCell>
            <TableCell align="center">เวลาที่บันทึก</TableCell>
            <TableCell align="center">Manage</TableCell>
         </TableRow>
       </TableHead>
       <TableBody>
         {predicament === undefined
            ? null
            : predicament.map((item) => (
           <TableRow key={item.id}>
             <TableCell align="center">{item.id}</TableCell>
             <TableCell align="center">{item.edges?.car?.carNo}</TableCell>
             <TableCell align="center">{item.edges?.staff?.staffName}</TableCell>
             <TableCell align="center">{item.edges?.statuscar?.statusDetail}</TableCell>
             <TableCell align="center">{item.edges?.user?.userName}</TableCell>
             <TableCell align="center">{moment(item.addedTime).format('yyyy-MM-ddThh:mm')}</TableCell>
             <TableCell align="center">
               <Button
                 onClick={() => {
                   if (item.id === undefined){
                     return;
                   }
                   deletePredicaments(item.id);
                 }}
                 style={{ marginLeft: 10 }}
                 variant="contained"
                 color="secondary"
               >
                 Delete
               </Button>
             </TableCell>
           </TableRow>
         ))}
       </TableBody>
     </Table>
   </TableContainer>
 );
}
