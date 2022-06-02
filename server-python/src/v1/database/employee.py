from src.v1.database.initial_db import SQLiteDatabase

class EmployeeQuery(SQLiteDatabase):
    def __init__(self, DBName) -> None:
        SQLiteDatabase.__init__(self, DBName)
        self.__connector = self.getConnector()
        self.__cur = self.getCursor()
        
    def get_all_employee(self):
        rows = self.__cur.execute(QueryString.get_all_employee())
        col = []
        for c in self.__cur.description:
            col.append(ConvertKey.convert_column(c[0]))
        return [dict(zip(col, row)) for row in rows.fetchall()]

class QueryString:
    @staticmethod
    def get_all_employee():
        return """SELECT * FROM employee"""
    
class ConvertKey:
    columns = {
        'id': 'ID',
        'first_name': 'FirstName',
        'last_name': 'LastName',
        'birth_date': 'BirthDate',
        'start_date': 'StartDate',
        'salary': 'Salary',
        'pvd_fund_rate': 'PvdFundRate',
    }
    
    @staticmethod
    def convert_column(col):
        if ConvertKey.columns.get(col, None):
            return ConvertKey.columns[col]
        else:
            return col
    
Qeury = EmployeeQuery('db.db')