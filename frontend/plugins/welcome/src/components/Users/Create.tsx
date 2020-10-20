import React, { useState, useEffect, FC} from 'react';
import { Link as RouterLink } from 'react-router-dom';
import SaveIcon from '@material-ui/icons/Save'; // icon save
import { Alert } from '@material-ui/lab';
import {
 Content,
 Header,
 Page,
 pageTheme,
 ContentHeader,
} from '@backstage/core';
import {
  Container,
  Grid,
  FormControl,
  Select,
  InputLabel,
  MenuItem,
  TextField,
  Avatar,
  Button,
  Link,
} from '@material-ui/core';
import { makeStyles, Theme, createStyles } from '@material-ui/core/styles';
import { DefaultApi } from '../../api/apis';
import Autocomplete from '@material-ui/lab/Autocomplete';
import { EntUser } from '../../api/models/EntUser';
import { EntStaff } from '../../api/models/EntStaff';
import { EntStatuscar } from '../../api/models/EntStatuscar';
import { EntCar } from '../../api/models/EntCar';
import { EntPredicament } from '../../api';

// css style 
const useStyles = makeStyles(theme => ({
  root: {
    flexGrow: 1,
  },
  paper: {
    marginTop: theme.spacing(2),
    marginBottom: theme.spacing(1),
  },
  formControl: {
    width: 300,
  },
  selectEmpty: {
    marginTop: theme.spacing(2),
  },
  container: {
    display: 'flex',
    flexWrap: 'wrap',
  },
  textField: {
    width: 300,
  },
  datetimepicker:{
    width: 300,
  },
}));
 
const HeaderCustom = {
  minHeight: '50px',
};

export default function Create() {
  const classes = useStyles();
  const api = new DefaultApi();

  const [predicament, setPredicament] = useState<EntPredicament>({});
  const [users, setUsers] = useState<EntUser[]>([]);
  const [cars, setCars] = useState<EntCar[]>([]);
  const [statuscars, setStatuscars] = useState<EntStatuscar[]>([]);
  const [staffs, setStaffs] = useState<EntStaff[]>([]);
  const [status, setStatus] = useState(false);
  const [alert, setAlert] = useState(true);
  const [loading, setLoading] = useState(true);
 
 
  const [userid, setUserid] = useState(Number);
  const [carid, setCarid] = useState(Number);
  const [staffid, setStaffid] = useState(Number);
  const [statuscarid, setStatuscarid] = useState(Number);
  const [addedTime, setAddedTime] = useState(String);
 
  useEffect(() => {
 
    const getCars = async () => {
 
      const cn = await api.listCar({ limit: 10, offset: 0 });
      setLoading(false);
      setCars(cn);
    };
    getCars();
 
    const getUsers = async () => {
 
    const u = await api.listUser({ limit: 10, offset: 0 });
      setLoading(false);
      setUsers(u);
    };
    getUsers();
 
    const getStatuscars = async () => {
 
     const s = await api.listStatuscar({ limit: 10, offset: 0 });
       setLoading(false);
       setStatuscars(s);
     };
     getStatuscars();

     const getStaffs = async () => {
 
      const st = await api.listStaff({ limit: 10, offset: 0 });
        setLoading(false);
        setStaffs(st);
      };
      getStaffs();
 
  }, [loading]);
 
  const CreatePredicament = async () => {
      const predicament = {
         addedTime      : addedTime + ":00+07:00", //+ "T00:00:00+07:00", //2020-10-20T11:53  yyyy-MM-ddT07:mm
         carID          : carid,
         staffID        : staffid,
         statuscarID    : statuscarid,
         userID         : userid,
      }
      console.log(predicament);

    const res:any = await api.createPredicament({ predicament : predicament});
    setStatus(true);
    if (res.id != ''){
      setAlert(true);
    } else {
      setAlert(false);
    }
 
    const timer = setTimeout(() => {
      setStatus(false);
    }, 1000);
  };
 
  const car_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
   setCarid(event.target.value as number);
    };

   const staff_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
     setStaffid(event.target.value as number);
    };

    const statuscar_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setStatuscarid(event.target.value as number);
     };

     const user_id_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setUserid(event.target.value as number);
     };
     const added_time_handleChange = (event: React.ChangeEvent<{ value: unknown }>) => {
      setAddedTime(event.target.value as string);
    };

 
  return (
    <Page theme={pageTheme.app}>
      <Header style={HeaderCustom} title={`ระบบการจัดการรถพยาบาล`}>
        <Avatar alt="Teerapat Saiprom" src="../../image/account.jpg" />
        <div style={{ marginLeft: 10 }}>Teerapat Saiprom</div>
      </Header>
      <Content>
      <ContentHeader title="บันทึกสถานะรถพยาบาล">
      <Link component={RouterLink} to="/results">
           <Button variant="contained" color="primary">
             ผลการบันทึก
           </Button>
         </Link>
         {status ? (
           <div>
             {alert ? (
               <Alert severity="success">
                 This is a success alert — check it out!
               </Alert>
             ) : (
               <Alert severity="warning" style={{ marginTop: 20 }}>
                 This is a warning alert — check it out!
               </Alert>
             )}
           </div>
         ) : null}
      </ContentHeader>
        <Container maxWidth="sm">
          <Grid container spacing={3}>
            <Grid item xs={12}></Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>หมายเลขรถพยาบาล</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกหมายเลขรถพยาบาล</InputLabel>
                <Select
                  name="car"
                  value={carid || ''}
                  onChange={car_id_handleChange}
                >
                  {cars.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.carNo}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>เจ้าหน้าที่ประจำรถพยาบาล</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกเจ้าหน้าที่ประจำรถพยาบาล</InputLabel>
                <Select
                  name="staff"
                  value={staffid || ''}
                  onChange={staff_id_handleChange}
                >
                  {staffs.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.staffName}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>สถานะรถพยาบาล</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกสถานะรถพยาบาล</InputLabel>
                <Select
                  name="statuscar"
                  value={statuscarid || ''}
                  onChange={statuscar_id_handleChange}
                >
                  {statuscars.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.statusDetail}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>

            <Grid item xs={4}>
              <div className={classes.paper}>เจ้าหน้าที่จัดการรถพยาบาล</div>
            </Grid>
            <Grid item xs={8}>
              <FormControl variant="outlined" className={classes.formControl}>
                <InputLabel>เลือกเจ้าหน้าที่จัดการรถพยาบาล</InputLabel>
                <Select
                  value={userid || ''}
                  onChange={user_id_handleChange}
                  name="user"
                >
                  {users.map(item => {
                    return (
                      <MenuItem key={item.id} value={item.id}>
                        {item.userEmail}
                      </MenuItem>
                    );
                  })}
                </Select>
              </FormControl>
            </Grid>
            <Grid item xs={4}>
              <div className={classes.paper}>เวลา</div>
            </Grid>
            <Grid item xs={8}>
            <form className={classes.container} noValidate>
                <TextField
                  label="เลือกเวลา"
                  name="added"
                  type="datetime-local"
                  value={addedTime || ''} // (undefined || '') = ''
                  className={classes.textField}
                  InputLabelProps={{
                    shrink: true,
                  }}
                  onChange={added_time_handleChange}
                />
              </form>
            </Grid>

            <Grid item xs={4}></Grid>
            <Grid item xs={8}>
              <Button
                variant="contained"
                color="primary"
                size="large"
                startIcon={<SaveIcon />}
                onClick={() => {
                  CreatePredicament();
                }}
              >
                บันทึกสถานะรถพยาบาล
              </Button>
            </Grid>
          </Grid>
        </Container>
      </Content>
    </Page>
  );
};
